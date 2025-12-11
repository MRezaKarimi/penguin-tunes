<script lang="ts" setup>
import SidebarItem from "./SidebarItem.vue";
import {
  IconHome,
  IconPlaylist,
  IconDisc,
  IconMusic,
  IconFolder,
  IconMicrophone2,
  IconSearch,
} from "@tabler/icons-vue";
</script>

<template>
  <div class="flex flex-col gap-1 shrink-0 w-64 bg-sidebar-bg px-4 py-5">
    <div class="flex items-center gap-3 py-1 px-2 bg-input rounded-lg mb-4">
      <IconSearch stroke="1.5" size="20" class="shrink-0" />
      <input
        class="outline-none border-none bg-transparent"
        placeholder="Search..."
        name="search"
      />
    </div>

    <SidebarItem>
      <IconHome stroke="1.25" />
      Home
    </SidebarItem>

    <SidebarItem>
      <IconPlaylist stroke="1.25" />
      Playlists
    </SidebarItem>

    <SidebarItem>
      <IconFolder stroke="1.25" />
      Folders
    </SidebarItem>

    <SidebarItem>
      <IconDisc stroke="1.25" />
      Albums
    </SidebarItem>

    <SidebarItem>
      <IconMusic stroke="1.25" />
      Tracks
    </SidebarItem>

    <SidebarItem>
      <IconMicrophone2 stroke="1.25" />
      Artists
    </SidebarItem>
  </div>
</template>

<style></style>
