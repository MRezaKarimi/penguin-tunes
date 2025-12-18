<script setup lang="ts">
import { computed } from "vue";
import { useMusicLibraryStore } from "@/lib/stores/library";
import {
  IconPlaylist,
  IconDisc,
  IconMusic,
  IconFolder,
  IconMicrophone2,
  IconPlayerPlayFilled,
} from "@tabler/icons-vue";
import { coverPathToURL } from "@/lib/utils";

const props = withDefaults(
  defineProps<{
    data: any;
    kind: "group" | "track";
  }>(),
  { kind: "group" }
);

const library = useMusicLibraryStore();

function getIcon() {
  if (library.view === "folder") {
    return IconFolder;
  } else if (library.view === "album") {
    return IconDisc;
  } else if (library.view === "track") {
    return IconMicrophone2;
  } else if (library.view === "playlist") {
    return IconPlaylist;
  } else {
    return IconMusic;
  }
}

const coverUrl = computed(() => {
  if (props.kind === "group") {
    const group = props.data;
    if (group?.tracks?.length > 0 && group.tracks[0].cover) {
      const firstTrackWithCover = group.tracks.find((t: any) => t.cover);
      return coverPathToURL(firstTrackWithCover?.cover);
    }
    return "";
  }

  return coverPathToURL(props.data.cover);
});

const title = computed(() => {
  if (props.kind === "group") {
    return props.data.name || "Unknown";
  }
  const t = props.data;
  return t?.title || t?.name || "Unknown Track";
});

const subtitle = computed(() => {
  if (props.kind === "group") {
    const g = props.data;
    return `${(g.tracks || []).length} songs`;
  }
  const t = props.data;
  return t?.artist || "Unknown Artist";
});
</script>

<template>
  <div class="relative flex flex-col cursor-pointer group">
    <div
      v-if="coverUrl"
      class="w-full aspect-square bg-cover bg-center"
      :style="{ 'background-image': `url('${coverUrl}')` }"
    />

    <div
      v-else
      class="flex flex-col justify-center items-center w-full aspect-square bg-cover bg-center bg-hover"
    >
      <component :is="getIcon()" size="50" stroke="1.25" />
    </div>

    <div class="text-sm truncate mt-2 mb-1">{{ title }}</div>
    <div class="text-xs font-light text-muted truncate">{{ subtitle }}</div>

    <div
      class="absolute top-1/2 end-2.5 bg-rose-500 rounded-full p-2 hover:scale-110 opacity-0 group-hover:opacity-100 transition-all"
    >
      <IconPlayerPlayFilled size="20" stroke="1.5" class="text-white" />
    </div>
  </div>
</template>
