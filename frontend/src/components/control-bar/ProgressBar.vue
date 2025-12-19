<script setup lang="ts">
import { computed } from "vue";
import { useQueueStore } from "@/lib/stores/queue";
import { formatTime } from "@/lib/utils";

const queue = useQueueStore();

const percent = computed(() => {
  const d = queue.duration || 0;
  if (!d) return 0;
  return Math.min(100, Math.max(0, (queue.currentTime / d) * 100));
});

function onClick(e: MouseEvent) {
  const target = e.currentTarget as HTMLElement;
  const rect = target.getBoundingClientRect();
  const x = e.clientX - rect.left;
  const pct = Math.max(0, Math.min(1, x / rect.width));
  const time = (queue.duration || 0) * pct;
  queue.seekTo(time);
}
</script>

<template>
  <div
    class="h-1.5 w-full bg-gray-300/40 rounded-full relative cursor-pointer group"
    @click="onClick"
  >
    <div
      id="seekFill"
      class="h-1.5 bg-rose-500 rounded-full transition-all duration-150"
      :style="{ width: percent + '%' }"
    />

    <div class="absolute -top-6 start-0 text-xs font-light">
      {{ formatTime(queue.currentTime) }}
    </div>
    <div class="absolute -top-6 end-0 text-xs font-light">
      {{ formatTime(queue.duration) }}
    </div>
  </div>
</template>
