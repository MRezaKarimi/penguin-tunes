export function coverPathToURL(path: string | undefined): string {
  if (path) {
    // keep basename so URL is /covers/<file>
    const coverFileName = path.split("/covers/").pop();
    return encodeURI(`/covers/${coverFileName}`);
  }
  return "";
}
