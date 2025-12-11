<template>
  <div
    class="h-1 w-full bg-gray-300/40 rounded-full relative cursor-pointer group"
  >
    <div
      id="seekFill"
      class="h-1 bg-rose-500 rounded-full transition-all duration-150"
      style="width: 30%"
    />
    <div
      id="seekKnob"
      class="absolute top-1/2 -translate-y-1/2 h-4 w-4 bg-rose-500 rounded-full shadow-lg opacity-0 group-hover:opacity-100 transition-opacity duration-150"
      style="left: 30%"
    />
  </div>
</template>
