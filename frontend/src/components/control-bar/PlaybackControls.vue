<script setup lang="ts">
import { useQueueStore } from "@/lib/stores/queue";
import {
  IconPlayerPlayFilled,
  IconPlayerPauseFilled,
  IconPlayerSkipForwardFilled,
  IconPlayerSkipBackFilled,
} from "@tabler/icons-vue";

const queue = useQueueStore();
</script>

<template>
  <div class="flex justify-center items-center gap-8 mb-4">
    <IconPlayerSkipBackFilled
      size="18"
      class="cursor-pointer"
      :class="{ 'pointer-events-none opacity-50': !queue.hasPrevious }"
      @click="queue.playPrevious"
    />
    <component
      :is="queue.isPlaying ? IconPlayerPauseFilled : IconPlayerPlayFilled"
      size="25"
      class="cursor-pointer"
      @click="queue.playPause"
    />
    <IconPlayerSkipForwardFilled
      size="18"
      class="cursor-pointer"
      :class="{ 'pointer-events-none opacity-50': !queue.hasNext }"
      @click="queue.playNext"
    />
  </div>
</template>
