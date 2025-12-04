<script lang="ts" setup>
import SidebarItem from "./SidebarItem.vue";
import {
  IconHome,
  IconPlaylist,
  IconDisc,
  IconMusic,
  IconFolder,
  IconMicrophone2,
} from "@tabler/icons-vue";
</script>

<template>
  <div class="flex flex-col gap-1 shrink-0 w-64 bg-sidebar-bg px-4 py-10">
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
