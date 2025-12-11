<template>
  <div class="flex items-center gap-4">
    <div class="flex items-center justify-center size-14 bg-gray-400">
      <IconDisc stroke="1.5" size="44" />
    </div>
    <div class="flex flex-col items-start gap-0.5">
      <span class="font-light">Blank Space</span>
      <span class="font-light cursor-pointer hover:underline text-xs"
        >1989</span
      >
      <span class="font-light cursor-pointer hover:underline text-xs"
        >Taylor Swift</span
      >
    </div>
  </div>
</template>

<script setup lang="ts">
import { IconDisc } from "@tabler/icons-vue";
</script>
