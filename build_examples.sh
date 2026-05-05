#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")" && pwd)"
BUILD_DIR="$ROOT_DIR/build"
EXAMPLES_DIR="$ROOT_DIR/examples"

mkdir -p "$BUILD_DIR"

WASM_FILES=()

for dir in "$EXAMPLES_DIR"/*/; do
  name="$(basename "$dir")"

  echo "==> Building $name..."

  # Change to the example directory so the build tool can resolve the package
  (cd "$dir" && go tool github.com/btwiuse/boba/cmd/boba-wasm-build \
    -o "$BUILD_DIR/$name.wasm" \
    .)

  echo "    -> $BUILD_DIR/$name.wasm"
  WASM_FILES+=("$name")
done

echo ""
echo "Done — all examples built into $BUILD_DIR"

# Generate index.html
INDEX="$BUILD_DIR/index.html"

cat > "$INDEX" <<EOF
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>bubbletea examples</title>
  <style>
    body { font-family: system-ui, sans-serif; max-width: 640px; margin: 40px auto; padding: 0 16px; }
    h1 { font-size: 1.5rem; }
    ul { padding: 0; list-style: none; }
    li { margin: 8px 0; }
    a { color: #0066cc; text-decoration: none; }
    a:hover { text-decoration: underline; }
  </style>
</head>
<body>
  <h1>bubbletea examples</h1>
  <ul>
EOF

for name in "${WASM_FILES[@]}"; do
  echo "    <li><a href=\"wasm.html?wasm=$name.wasm\">$name</a></li>" >> "$INDEX"
done

cat >> "$INDEX" <<EOF
  </ul>
</body>
</html>
EOF

echo "    -> $INDEX"

echo "bundling ./build/wasm.html"
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" ./assets
bun build --compile --outfile=./build/wasm.html --target=browser ./assets/index.html
