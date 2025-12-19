<script lang="ts" setup>
import { useMusicLibraryStore } from "@/lib/stores/library";
import Header from "@/components/Header.vue";
import Item from "./Item.vue";

const library = useMusicLibraryStore();
</script>

<template>
  <div class="flex flex-col grow relative">
    <Header />
    <div
      class="grow overflow-auto p-4 gap-x-5 gap-y-8"
      style="
        display: grid;
        grid-template-columns: repeat(
          auto-fit,
          minmax(clamp(90px, 16vw, 8rem), 1fr)
        );
        justify-content: start;
        align-content: start;
      "
    >
      <Item v-for="item in library.tracks" :key="item.id" :data="item" />

      <!-- Spacer divs to ensure proper grid layout -->
      <div v-for="i in 10" :key="i" />
    </div>
  </div>
</template>
