<script setup lang="ts">
import { computed } from "vue";
import { useMusicLibraryStore } from "@/lib/stores/library";
import { coverPathToURL } from "@/lib/utils";
import { useQueueStore } from "@/lib/stores/queue";
import {
  IconPlaylist,
  IconDisc,
  IconMusic,
  IconFolder,
  IconMicrophone2,
  IconPlayerPlayFilled,
} from "@tabler/icons-vue";

const library = useMusicLibraryStore();
const queue = useQueueStore();

const props = defineProps<{ data: (typeof library.tracks)[number] }>();

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
  if ("tracks" in props.data) {
    const group = props.data;
    if (group?.tracks?.length > 0 && group.tracks[0].cover) {
      const firstTrackWithCover = group.tracks.find((t: any) => t.cover);
      return coverPathToURL(firstTrackWithCover?.cover);
    }
    return "";
  }

  return coverPathToURL(props.data.cover);
});

const subtitle = computed(() => {
  if ("tracks" in props.data) {
    const g = props.data;
    return `${(g.tracks || []).length} songs`;
  }
  const t = props.data;
  return t?.artist || "Unknown Artist";
});
</script>

<template>
  <div
    class="relative flex flex-col cursor-pointer group"
    @click="'path' in data && queue.playTrack(data)"
  >
    <!-- Cover image if exists -->
    <div
      v-if="coverUrl"
      class="w-full aspect-square bg-cover bg-center"
      :style="{ 'background-image': `url('${coverUrl}')` }"
    />
    <!-- Otherwise show an icon as a placeholder -->
    <div
      v-else
      class="flex flex-col justify-center items-center w-full aspect-square bg-cover bg-center bg-hover"
    >
      <component :is="getIcon()" size="50" stroke="1.25" />
    </div>

    <div class="text-sm truncate mt-2 mb-1">{{ data.title }}</div>
    <div class="text-xs font-light text-muted truncate">{{ subtitle }}</div>

    <!-- Add current group to the queue and play the first track in it -->
    <div
      v-if="'tracks' in data"
      class="absolute top-1/2 end-2.5 bg-rose-500 rounded-full p-2 hover:scale-110 opacity-0 group-hover:opacity-100 transition-all"
      @click.stop="queue.playGroup(data.tracks)"
    >
      <IconPlayerPlayFilled size="20" stroke="1.5" class="text-white" />
    </div>
  </div>
</template>
