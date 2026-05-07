const loading = document.getElementById("loading");

function waitForBridge() {
  return new Promise((resolve) => {
    function check() {
      if (
        globalThis.bubbletea_resize !== undefined &&
        globalThis.bubbletea_read !== undefined &&
        globalThis.bubbletea_write !== undefined
      ) {
        resolve();
      } else {
        console.log("waiting for bubbletea bridge…");
        setTimeout(check, 500);
      }
    }
    check();
  });
}

function initTerminal() {
  const term = new Terminal({
    convertEol: true,
    cursorBlink: true,
    allowTransparency: true,
  });
  const imageAddon = new ImageAddon.ImageAddon();
  term.loadAddon(imageAddon);
  const fitAddon = new FitAddon.FitAddon();
  if (new URLSearchParams(location.search).get("webgl") !== null) {
    const webglAddon = new WebglAddon.WebglAddon();
    try {
      term.loadAddon(webglAddon);
    } catch (e) {
      console.warn(
        "WebGL addon failed to load, falling back to canvas renderer",
        e,
      );
    }
  }
  term.loadAddon(fitAddon);
  term.open(document.getElementById("terminal-container"));

  fitAddon.fit();
  window.addEventListener("resize", () => fitAddon.fit());

  term.focus();

  // Send initial size to Go
  bubbletea_resize(term.cols, term.rows);

  /** Whether the Go program has exited; gate all input after this point. */
  let exited = false;

  // Poll Go output and write to terminal
  const pollInterval = setInterval(() => {
    if (exited) return;
    const data = bubbletea_read();
    if (data && data.length > 0) {
      term.write(data);
    }
  }, 16);

  // Forward resize events to Go
  term.onResize((size) => {
    if (!exited) bubbletea_resize(size.cols, size.rows);
  });

  // Forward key/paste input to Go; reload after exit
  term.onData((data) => {
    if (exited) {
      location.reload();
      return;
    }
    bubbletea_write(data);
  });

  return {
    term,
    pollInterval,
    setExited: (v) => {
      exited = v;
    },
  };
}

async function main() {
  const go = new Go();
  const wasmPath = new URLSearchParams(location.search).get("wasm") ||
    "./booba.wasm";
  const result = await WebAssembly.instantiateStreaming(
    fetch(wasmPath),
    go.importObject,
  );

  // Start the WASM module (non-blocking); Go registers the bridge globals as it runs
  const runPromise = go.run(result.instance);

  // Wait until go-booba registers the JS bridge globals
  await waitForBridge();

  // Hide the loading overlay
  loading.classList.add("hidden");

  const { term, pollInterval, setExited } = initTerminal();

  // When the Go program exits, show a restart prompt
  runPromise.then(() => {
    console.log("wasm finished");
    setExited(true);
    clearInterval(pollInterval);
    term.write("\r\n\r\nPress any key to continue...");
  });
}

main().catch(console.error);
