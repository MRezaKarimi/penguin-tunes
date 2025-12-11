<script setup lang="ts">
import { IconDisc } from "@tabler/icons-vue";
import { useTrackStore } from "@/lib/stores/track";

const track = useTrackStore();
</script>

<template>
  <div class="flex items-center gap-4">
    <div class="flex items-center justify-center size-14 bg-gray-400">
      <IconDisc stroke="1.5" size="44" />
    </div>
    <div class="flex flex-col items-start gap-0.5">
      <span class="font-light">{{ track.name }}</span>
      <span class="font-light cursor-pointer hover:underline text-xs">{{
        track.album
      }}</span>
      <span class="font-light cursor-pointer hover:underline text-xs">{{
        track.artist
      }}</span>
    </div>
  </div>
</template>
