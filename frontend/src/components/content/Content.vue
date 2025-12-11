<script lang="ts" setup>
import Header from "../Header.vue";
import Item from "./Item.vue";
</script>

<template>
  <div class="flex flex-col grow relative">
    <Header />

    <div
      class="grow overflow-auto p-4"
      style="
        display: grid;
        grid-template-columns: repeat(
          auto-fit,
          minmax(clamp(100px, 18vw, 150px), 1fr)
        );
        gap: 1.25rem;
        align-content: start;
      "
    >
      <Item v-for="i in 40" :key="i" />
    </div>
  </div>
</template>
