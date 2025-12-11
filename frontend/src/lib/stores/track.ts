import { defineStore } from "pinia";
import { ref } from "vue";

export const useTrackStore = defineStore("track", () => {
  const name = ref("");
  const album = ref("");
  const artist = ref("");

  function setTrackInfo(track: {
    name: string;
    album: string;
    artist: string;
  }) {
    name.value = track.name;
    album.value = track.album;
    artist.value = track.artist;
  }

  return { name, album, artist, setTrackInfo };
});
