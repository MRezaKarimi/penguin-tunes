<script setup lang="ts">
import {
  IconPlayerPlayFilled,
  IconPlayerSkipForwardFilled,
  IconPlayerSkipBackFilled,
} from "@tabler/icons-vue";
</script>

<template>
  <div class="flex justify-center items-center gap-8 mb-4">
    <IconPlayerSkipBackFilled size="18" class="cursor-pointer" />
    <IconPlayerPlayFilled size="25" class="cursor-pointer" />
    <IconPlayerSkipForwardFilled size="18" class="cursor-pointer" />
  </div>
</template>
