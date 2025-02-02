<script setup lang="ts">
import RelatedChannels from '@/components/channel/RelatedChannels.vue';

const route = useRoute();
const channelId = computed(() => getChannelIdFromParam(route.params.id));
const { data: channelInfo, pending } = useGetChannelInfo(channelId);
const { data: channelHome, pending: pendingHome } = useGetChannelHome(channelId);
</script>

<template>
  <div v-if="!pending && !pendingHome && channelInfo && channelHome" class="channel-home">
    <SectionTitle title="Info" />
    <pre v-if="channelInfo.description" class="channel-description">{{
      channelInfo.description?.trim()
    }}</pre>
    <ChannelBannerLinks
      v-if="channelInfo.channelLinks"
      :banner-links="{ ...channelInfo?.channelLinks, type: 'links' }"
    />
    <SectionTitle v-if="channelInfo.relatedChannels?.items?.length > 0" title="Related channels" />
    <RelatedChannels
      v-if="channelInfo.relatedChannels?.items?.length > 0"
      :related-channels="{ ...channelInfo.relatedChannels, type: 'channels' }"
    />
    <SectionTitle v-if="channelHome.featuredVideo" title="Featured video" />
    <ChannelFeaturedVideo
      v-if="channelHome.featuredVideo"
      :featured-video="channelHome.featuredVideo"
    />
    <div v-for="(shelf, index) in channelHome?.items ?? []" :key="index" class="shelves">
      <SectionTitle :title="shelf.shelfName" :link="shelf.shelfUrl" />
      <ChannelPlaylistShelf
        v-if="shelf.type === 'playlist' || shelf.type === 'videos'"
        :shelf="shelf"
      />
      <ChannelPlaylistsShelf v-else-if="shelf.type === 'playlists'" :shelf="shelf" />
      <ChannelFeaturedChannelsShelf v-else-if="shelf.type === 'channels'" :shelf="shelf" />
    </div>
  </div>
</template>

<style lang="scss" scoped>
.channel-home {
  padding: 0 10px;

  .channel-description {
    white-space: pre-wrap;
    font-family: $default-font;
    margin: 0;
    word-break: break-word;
  }
}
</style>
