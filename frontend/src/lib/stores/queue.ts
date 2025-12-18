import { defineStore } from "pinia";
import { computed, ref } from "vue";

export const useQueueStore = defineStore("queue", () => {
  const queue = ref<any[]>([]);
  const index = ref(0);
  const isPlaying = ref(false);

  const currentlyPlaying = computed(() => {
    return queue.value[index.value];
  });

  function playNext() {
    if (queue.value[index.value + 1]) {
      index.value += 1;
    }
  }

  function playPrevious() {
    if (queue.value[index.value - 1]) {
      index.value -= 1;
    }
  }

  function addToQueue(item: any) {
    queue.value.push(item);
  }

  function clearQueue() {
    queue.value = [];
  }

  function setQueue(newQueue: any[], startIndex: number = 0) {
    queue.value = newQueue;
    index.value = startIndex;
  }

  function playPause() {}

  function seekTo(time: number) {}

  return {
    isPlaying,
    index,
    currentlyPlaying,
    playNext,
    playPrevious,
    addToQueue,
    setQueue,
    clearQueue,
    playPause,
    seekTo,
  };
});
