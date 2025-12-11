<script setup lang="ts">
import {
  IconPlaylistAdd,
  IconVolume,
  IconAdjustments,
  IconArrowsShuffle,
} from "@tabler/icons-vue";
</script>

<template>
  <div class="flex items-center gap-4 pe-8">
    <IconPlaylistAdd stroke="1.75" size="24" class="cursor-pointer" />
    <IconVolume stroke="1.75" size="20" class="cursor-pointer" />
    <IconAdjustments stroke="1.75" size="20" class="cursor-pointer" />
    <IconArrowsShuffle stroke="1.75" size="20" class="cursor-pointer" />
  </div>
</template>
