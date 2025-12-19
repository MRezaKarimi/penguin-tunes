import { Track } from "@/types";

export function coverPathToURL(path: string | undefined): string {
  if (path) {
    // keep basename so URL is /covers/<file>
    const coverFileName = path.split("/covers/").pop();
    return encodeURI(`/covers/${coverFileName}`);
  }
  return "";
}

export function makeSrcForTrack(t?: Track | null) {
  if (!t) return null;
  // support both `path` and `Path` variants
  const p: string | undefined = (t as any).path || (t as any).Path;
  if (!p) return null;

  // Prefer an explicit HTTP URL when running inside Wails so the request
  // reaches the Go backend instead of being treated as a Wails app asset.
  const { protocol, hostname, port } = window.location;
  let base: string;

  if (protocol && protocol.startsWith("wails")) {
    // wails://... origins don't resolve via normal DNS inside the webview,
    // so use localhost/127.0.0.1 and the dev port (34115) as a fallback.
    const host =
      hostname === "wails.localhost" || hostname.endsWith(".wails.localhost")
        ? "127.0.0.1"
        : hostname;
    const effectivePort = port || "34115";
    base = `http://${host}:${effectivePort}`;
  } else {
    // normal browsers: keep same origin (preserves host/port)
    base = `${protocol}//${hostname}${port ? `:${port}` : ""}`;
  }

  const src = `${base}/asset?path=${encodeURIComponent(p)}`;
  return src;
}

export function formatTime(seconds: number): string {
  const mins = Math.floor(seconds / 60);
  const secs = Math.floor(seconds % 60);
  return `${mins.toString().padStart(2, "0")}:${secs.toString().padStart(2, "0")}`;
}
