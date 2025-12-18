<script setup lang="ts">
import TrackInfo from "./TrackInfo.vue";
import PlaybackControls from "./PlaybackControls.vue";
import ProgressBar from "./ProgressBar.vue";
import AdditionalControls from "./AdditionalControls.vue";
</script>

<template>
  <div class="flex items-center p-4 border-t-2 border-sidebar-bg">
    <TrackInfo />

    <div class="w-full max-w-xl mx-auto">
      <PlaybackControls />
      <ProgressBar />
    </div>

    <AdditionalControls />
  </div>
</template>
