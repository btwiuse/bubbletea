# Bouncing DVD Logo — 实现计划

## 目标

在 terminal 中实现一个模仿 DVD 待机画面的弹跳 Logo demo，支持鼠标拖拽 Logo。

## 参考

`./clickable/` — 使用 `lipgloss.NewLayer` + `lipgloss.NewCompositor` 实现可拖拽 dialog。

## 结构

```
bouncing-dvd-logo/
├── main.go    # Model、Update、View、main
└── plan.md    # 本文件
```

只用一个文件（比 clickable 简单，无按钮/多 dialog 等）。

## 核心逻辑

### 1. Model

```go
type model struct {
    x, y          int      // 当前位置（左上角）
    vx, vy        int      // 速度方向（每帧像素），如 +1/+1 或 -1/+1
    width, height int      // terminal 大小
    dragging      bool     // 是否正在拖拽
    dragOffX      int      // 鼠标相对于 Logo 左上角的偏移
    dragOffY      int
    paused        bool     // 拖拽时暂停自动移动
}
```

### 2. Logo 外观

用带边框的方框 + "DVD" 文字。简单即可：

```
┌─────┐
│ DVD │
└─────┘
```

使用 `lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Width(w).Height(h)` 渲染。

Box 尺寸固定（如 width=9, height=5）。

支持多个颜色循环（每个弹跳换色），或者固定一种亮色。

### 3. 动画循环

使用 `tea.TickMsg` + 自定义 interval（如 50ms）驱动帧更新。

每帧：
- 如果不是拖拽暂停状态，则 `x += vx; y += vy`
- 撞到边界则反转对应方向的分量

### 4. 碰撞检测（Physics）

```
if x <= 0 || x + boxWidth >= termWidth  => vx = -vx
if y <= 0 || y + boxHeight >= termHeight => vy = -vy
```

撞墙后 clamp 位置到边界内，防止卡在墙外。

### 5. 鼠标拖拽交互

复用 clickable 的 `LayerHitMsg` 模式：

- `tea.MouseClickMsg`（左键按下）：检测是否点中 Logo 层。如果是：
  - 暂停动画
  - 记录偏移量 (`dragOffX = mouseX - logo.x`)
  - 置 `dragging = true`
- `tea.MouseMotionMsg`（移动中 + 按下）：如果 `dragging`：
  - `logo.x = mouseX - dragOffX`
  - `logo.y = mouseY - dragOffY`
  - clamp 到边界
- `tea.MouseReleaseMsg`：释放：
  - `dragging = false`
  - 恢复动画

释放 Logo 后，速度方向保持不变（继续沿当前方向弹跳）。

### 6. View

使用 `lipgloss.NewCompositor` + layers：

- 背景层（id="bg"），用 `/` 或空白填充
- Logo 层（id="logo"），带 border 的 box

`OnMouse` 回调通过 `comp.Hit(x, y)` 判断鼠标命中了哪一层，派发 `LayerHitMsg`。

### 7. Key 绑定

- `q` / `ctrl+c` / `esc` → quit

## 可选增强（先不做）

- 多个 DVD Logo
- 撞墙变色 + 变色计数
- 撞墙音效
- 速度逐渐加快
- TUI 启动画面风格（深色背景 + 彩色 Logo）

## 验证方式

```bash
cd examples/bouncing-dvd-logo
go run .
```

确认：
1. Logo 自动沿对角线弹跳，碰到边缘正确反弹
2. 鼠标按住 Logo 可拖拽，松手继续弹跳
3. Logo 不会超出 terminal 边界
4. 拖拽时动画暂停，不跳帧
