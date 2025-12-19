<script setup lang="ts">
import { IconDisc } from "@tabler/icons-vue";
import { useQueueStore } from "@/lib/stores/queue";
import { coverPathToURL } from "@/lib/utils";

const queue = useQueueStore();
</script>

<template>
  <div class="w-56 shrink-0">
    <div v-if="queue.currentlyPlaying" class="flex items-center gap-4">
      <!-- Cover art -->
      <div
        v-if="queue.currentlyPlaying.cover"
        class="size-14 shrink-0 aspect-square bg-cover bg-center"
        :style="{
          'background-image': `url('${coverPathToURL(queue.currentlyPlaying.cover)}')`,
        }"
      />
      <!-- Otherwise show an icon as a placeholder -->
      <div
        v-else
        class="flex items-center justify-center size-14 shrink-0 bg-gray-400"
      >
        <IconDisc stroke="1.5" size="44" />
      </div>

      <div class="flex flex-col items-start gap-0.5">
        <span class="font-light truncate">
          {{ queue.currentlyPlaying.title }}
        </span>
        <span
          class="font-light truncate cursor-pointer hover:underline text-xs"
        >
          {{ queue.currentlyPlaying.album }}
        </span>
        <span class="font-light cursor-pointer hover:underline text-xs">
          {{ queue.currentlyPlaying.artist }}
        </span>
      </div>
    </div>
  </div>
</template>
