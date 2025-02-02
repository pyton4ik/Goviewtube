import { ChannelCommunityPostsContinuationResponse } from '../../types/ytch.types';
import { ChannelCommunityPostDto } from '../basic/channel-community-post.dto';

export class ChannelCommunityPostsContinuationDto
  implements ChannelCommunityPostsContinuationResponse
{
  innerTubeApi: string;
  items: ChannelCommunityPostDto[];
  continuation: string;
}
