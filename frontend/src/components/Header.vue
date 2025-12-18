<script lang="ts" setup>
import {
  IconX,
  IconWindowMinimize,
  IconMaximize,
  IconArrowLeft,
  IconArrowRight,
} from "@tabler/icons-vue";
import { useTrackStore } from "@/lib/stores/track";

const track = useTrackStore();
</script>

<template>
  <header
    class="flex justify-center items-center gap-5 shrink-0 h-12 px-5"
    style="--wails-draggable: drag"
  >
    <!-- Back/Forward -->
    <span class="cursor-pointer" onclick="window.runtime.Quit()">
      <IconArrowLeft stroke="1.5" size="20" />
    </span>
    <span class="cursor-pointer" onclick="window.runtime.Quit()">
      <IconArrowRight stroke="1.5" size="20" />
    </span>

    <!-- Title -->
    <div
      class="grow select-none"
      ondblclick="window.runtime.WindowToggleMaximise()"
    >
      {{ track.name || "PenguinTunes" }}
    </div>

    <!-- Window Controls -->
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
