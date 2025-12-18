<script setup lang="ts">
defineProps<{ isActive?: boolean }>();
</script>
<template>
  <div
    :class="{ 'bg-hover': isActive }"
    class="flex items-center gap-2 rounded-lg cursor-pointer font-light hover:bg-hover py-2 px-3"
  >
    <slot />
  </div>
</template>
