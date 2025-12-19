import { defineStore } from "pinia";
import { computed, ref } from "vue";
import { Track } from "@/types";
import { makeSrcForTrack } from "../utils";

export const useQueueStore = defineStore("queue", () => {
  const queue = ref<Track[]>([]);
  const index = ref(0);
  const isPlaying = ref(false);
  const currentTime = ref(0); // seconds
  const duration = ref(0); // seconds (total time of current track)
  const currentSrc = ref<string | null>(null);

  // singleton audio element used by the store
  const audio = new Audio();
  audio.preload = "metadata";
  audio.crossOrigin = "anonymous";

  const currentlyPlaying = computed(() => {
    return queue.value[index.value];
  });

  const hasNext = computed(() => {
    return !!queue.value[index.value + 1];
  });

  const hasPrevious = computed(() => {
    return !!queue.value[index.value - 1];
  });

  function loadCurrentTrack(autoplay = true) {
    const t = currentlyPlaying.value;
    const src = makeSrcForTrack(t);
    if (!src) {
      audio.pause();
      currentSrc.value = null;
      isPlaying.value = false;
      currentTime.value = 0;
      duration.value = 0;
      audio.removeAttribute("src");
      return;
    }

    if (audio.src !== src) {
      audio.src = src;
      currentSrc.value = src;
      audio.load();
    }

    if (autoplay && isPlaying.value) {
      audio.play().catch(() => {
        // play may be blocked by autoplay policy; reflect paused state
        isPlaying.value = false;
      });
    }
  }

  // audio event wiring
  audio.addEventListener("loadedmetadata", () => {
    duration.value = isFinite(audio.duration) ? audio.duration : 0;
  });

  audio.addEventListener("timeupdate", () => {
    currentTime.value = audio.currentTime || 0;
  });

  audio.addEventListener("ended", () => {
    // automatically go to next track if available
    if (hasNext.value) {
      playNext();
    } else {
      isPlaying.value = false;
      currentTime.value = 0;
    }
  });

  audio.addEventListener("play", () => (isPlaying.value = true));
  audio.addEventListener("pause", () => (isPlaying.value = false));
  audio.addEventListener("error", () => {
    // on error, stop playback and try to advance
    console.warn("Audio playback error", audio.error);
    isPlaying.value = false;
  });

  function playNext() {
    if (queue.value[index.value + 1]) {
      index.value += 1;
      currentTime.value = 0;
      // load and autoplay the next track
      isPlaying.value = true;
      loadCurrentTrack(true);
    }
  }

  function playPrevious() {
    if (queue.value[index.value - 1]) {
      index.value -= 1;
      currentTime.value = 0;
      isPlaying.value = true;
      loadCurrentTrack(true);
    }
  }

  function clearQueue() {
    queue.value = [];
    index.value = 0;
    isPlaying.value = false;
    loadCurrentTrack(false);
  }

  function setQueue(newQueue: Track[], startIndex: number = 0) {
    queue.value = newQueue || [];
    index.value = Math.max(
      0,
      Math.min(startIndex || 0, queue.value.length - 1)
    );
    currentTime.value = 0;
    isPlaying.value = true;
    loadCurrentTrack(true);
  }

  function playTrack(track: Track) {
    setQueue([track], 0);
  }

  function playGroup(tracks: Track[], startIndex: number = 0) {
    setQueue(tracks, startIndex);
  }

  function setDuration(totalSeconds: number) {
    duration.value = totalSeconds;
  }

  function getProgress() {
    if (!duration.value) return 0;
    return Math.min(1, Math.max(0, currentTime.value / duration.value));
  }

  function togglePlayPause() {
    if (!currentlyPlaying.value) return;
    if (audio.paused) {
      // try to play
      audio.play().catch(() => {
        isPlaying.value = false;
      });
    } else {
      audio.pause();
    }
    // isPlaying will be updated by play/pause events
  }

  function seekTo(time: number) {
    if (!currentlyPlaying.value) return;
    // clamp
    const t = Math.max(0, Math.min(time || 0, duration.value || Infinity));
    try {
      audio.currentTime = t;
      currentTime.value = t;
    } catch (err) {
      // some browsers may throw if not ready; ignore
      console.warn("seek failed", err);
    }
  }

  return {
    queue,
    isPlaying,
    index,
    currentlyPlaying,
    currentTime,
    duration,
    currentSrc,
    hasNext,
    hasPrevious,
    getProgress,
    playNext,
    playPrevious,
    setQueue,
    clearQueue,
    playPause: togglePlayPause,
    seekTo,
    playTrack,
    playGroup,
    setDuration,
  };
});
