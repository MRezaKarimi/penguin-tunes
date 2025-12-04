<script lang="ts" setup>
import { IconX, IconWindowMinimize, IconMaximize } from "@tabler/icons-vue";
</script>

<template>
  <header
    class="flex justify-end items-center gap-5 w-full shrink-0 py-3 px-5"
    style="--wails-draggable: drag"
  >
    <span class="cursor-pointer" onclick="window.runtime.WindowMinimise()">
      <IconWindowMinimize stroke="1.5" size="18" />
    </span>
    <span
      class="cursor-pointer"
      onclick="window.runtime.WindowToggleMaximise()"
    >
      <IconMaximize stroke="1.5" size="18" />
    </span>
    <span class="cursor-pointer" onclick="window.runtime.Quit()">
      <IconX stroke="1.5" size="18" />
    </span>
  </header>
</template>
