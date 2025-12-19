import { computed, ref } from "vue";
import { defineStore } from "pinia";
import { Track } from "@/types";
import { ReadIndex } from "../../../wailsjs/go/main/App";
import * as runtime from "../../../wailsjs/runtime/runtime";

export const useMusicLibraryStore = defineStore("library", () => {
  const _tracks = ref<Track[]>([]);
  const view = ref<"artist" | "album" | "folder" | "track" | "playlist">(
    "artist"
  );

  const tracks = computed(() => {
    if (view.value === "track" || view.value === "playlist") {
      return _tracks.value;
    }

    if (view.value === "folder") {
      const groups: Record<string, Track[]> = {};
      for (const t of _tracks.value) {
        const path = t.path;
        const dir = path.substring(0, Math.max(0, path.lastIndexOf("/")));
        groups[dir] = groups[dir] || [];
        groups[dir].push(t);
      }
      return Object.keys(groups).map((k) => ({
        id: k,
        title: k.split("/").pop(),
        tracks: groups[k],
      }));
    }

    const groups: Record<string, Track[]> = {};
    for (const track of _tracks.value) {
      const key = track[view.value as keyof Track] || `Unknown ${view.value}`;
      groups[key] = groups[key] || [];
      groups[key].push(track);
    }
    return Object.keys(groups).map((k) => ({
      id: k,
      title: k,
      tracks: groups[k],
    }));
  });

  function setView(v: typeof view.value) {
    view.value = v;
  }

  async function loadTracks() {
    const rawIndex = (await ReadIndex()) || "";
    if (rawIndex) {
      _tracks.value = parseIndex(rawIndex);
    }
  }

  function parseIndex(rawIndex: string): Track[] {
    try {
      const parsed = JSON.parse(rawIndex);
      return parsed.tracks ? Object.values(parsed.tracks) : [];
    } catch (error) {
      console.error(error);
      return [];
    }
  }

  // Listen for index updates from the backend
  runtime.EventsOn("index-updated", () => {
    loadTracks();
  });

  loadTracks();

  return {
    tracks,
    view,
    setView,
  };
});
