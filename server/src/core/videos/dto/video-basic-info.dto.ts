import { AuthorThumbnailDto } from './author-thumbnail.dto';
import { VideoThumbnailDto } from './video-thumbnail.dto';

export class VideoBasicInfoDto {
  videoId: string;
  title: string;
  published?: number;
  publishedText: string;
  author: string;
  authorId: string;
  authorVerified?: boolean;
  authorThumbnails?: Array<AuthorThumbnailDto>;
  authorThumbnailUrl?: string;
  videoThumbnails: Array<VideoThumbnailDto>;
  description?: string;
  viewCount: number;
  likeCount?: number;
  dislikeCount?: number;
  lengthSeconds?: number;
  lengthString?: string;
  live?: boolean;
}
