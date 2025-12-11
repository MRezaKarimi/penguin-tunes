<script lang="ts" setup>
import { IconX, IconWindowMinimize, IconMaximize } from "@tabler/icons-vue";
</script>

<template>
  <header
    class="relative flex justify-center items-center shrink-0 h-12"
    style="--wails-draggable: drag"
  >
    <div
      class="absolute end-0 flex justify-end items-center gap-5 w-full shrink-0 py-3 px-5"
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
    </div>

    <div>Blank Space</div>
  </header>
</template>
