import { ChannelCommunityPostsResponse } from '../../types/ytch.types';
import { ChannelCommunityPostDto } from '../basic/channel-community-post.dto';

export class ChannelCommunityPostsDto implements ChannelCommunityPostsResponse {
  channelIdType: number;
  innerTubeApi: string;
  items: ChannelCommunityPostDto[];
  continuation: string;
}
