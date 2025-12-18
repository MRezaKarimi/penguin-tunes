<script setup lang="ts">
import { IconDisc } from "@tabler/icons-vue";
import { useQueueStore } from "@/lib/stores/queue";

const queue = useQueueStore();
</script>

<template>
  <div class="w-56 shrink-0">
    <div v-if="queue.currentlyPlaying" class="flex items-center gap-4">
      <div class="flex items-center justify-center size-14 bg-gray-400">
        <IconDisc stroke="1.5" size="44" />
      </div>
      <div class="flex flex-col items-start gap-0.5">
        <span class="font-light">{{ queue.currentlyPlaying.title }}</span>
        <span class="font-light cursor-pointer hover:underline text-xs">
          {{ queue.currentlyPlaying.album }}
        </span>
        <span class="font-light cursor-pointer hover:underline text-xs">
          {{ queue.currentlyPlaying.artist }}
        </span>
      </div>
    </div>
  </div>
</template>
