<script lang="ts" setup>
import Content from "./components/content/Content.vue";
import ControlBar from "./components/control-bar/ControlBar.vue";
import Sidebar from "./components/sidebar/Sidebar.vue";
</script>

<template>
  <div class="flex flex-col h-full w-full">
    <div class="flex grow overflow-auto">
      <Sidebar />
      <Content />
    </div>
    <ControlBar />
  </div>
</template>

<style></style>
