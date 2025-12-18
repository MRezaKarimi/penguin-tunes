<script lang="ts" setup>
import { useMusicLibraryStore } from "@/lib/stores/library";
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

const library = useMusicLibraryStore();
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

    <SidebarItem
      :is-active="library.view === 'playlist'"
      @click="library.setView('playlist')"
    >
      <IconPlaylist stroke="1.25" />
      Playlists
    </SidebarItem>

    <SidebarItem
      :is-active="library.view === 'folder'"
      @click="library.setView('folder')"
    >
      <IconFolder stroke="1.25" />
      Folders
    </SidebarItem>

    <SidebarItem
      :is-active="library.view === 'album'"
      @click="library.setView('album')"
    >
      <IconDisc stroke="1.25" />
      Albums
    </SidebarItem>

    <SidebarItem
      :is-active="library.view === 'track'"
      @click="library.setView('track')"
    >
      <IconMusic stroke="1.25" />
      Tracks
    </SidebarItem>

    <SidebarItem
      :is-active="library.view === 'artist'"
      @click="library.setView('artist')"
    >
      <IconMicrophone2 stroke="1.25" />
      Artists
    </SidebarItem>
  </div>
</template>

<style></style>
