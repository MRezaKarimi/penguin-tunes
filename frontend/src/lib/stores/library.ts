import { ref } from "vue";
import { defineStore } from "pinia";
import {
  GetTracks,
  GetAlbums,
  GetArtists,
  GetFolders,
} from "../../../wailsjs/go/main/App";
import * as runtime from "../../../wailsjs/runtime/runtime";

export const useMusicLibraryStore = defineStore("library", () => {
  const tracks = ref<any[]>([]);
  const view = ref<"artist" | "album" | "folder" | "track" | "playlist">(
    "artist"
  );

  async function loadByView(v: typeof view.value) {
    if (v === "artist") {
      tracks.value = (await GetArtists()) as any[];
    } else if (v === "album") {
      tracks.value = (await GetAlbums()) as any[];
    } else if (v === "folder") {
      tracks.value = (await GetFolders()) as any[];
    } else {
      tracks.value = (await GetTracks()) as any[];
    }
  }

  function setView(v: typeof view.value) {
    view.value = v;
    loadByView(v);
  }

  async function loadTracks() {
    await loadByView(view.value);
  }

  // Listen for index updates from the backend
  runtime.EventsOn("index-updated", () => {
    loadByView(view.value);
  });

  loadTracks();

  return {
    tracks,
    view,
    setView,
  };
});
