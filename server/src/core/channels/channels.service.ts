import path from 'path';
import fs from 'fs';
import { BadRequestException, Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { General } from 'server/common/general.schema';
import { FastifyReply } from 'fastify';
import sharp from 'sharp';
import ytch from 'yt-channel-info';
import { ChannelHomeDto } from './dto/response/channel-home.dto';
import { ChannelVideosDto } from './dto/response/channel-videos.dto';
import { ChannelVideosContinuationDto } from './dto/response/channel-videos-continuation.dto';
import { checkParams, throwChannelError } from './channels.helper';
import { ChannelPlaylistsDto } from './dto/response/channel-playlists.dto';
import { ChannelPlaylistsContinuationDto } from './dto/response/channel-playlists-continuation.dto';
import { ChannelSearchDto } from './dto/response/channel-search.dto';
import { ChannelSearchContinuationDto } from './dto/response/channel-search-continuation.dto';
import { RelatedChannelsContinuationDto } from './dto/response/related-channels-continuation.dto';
import { ChannelCommunityPostsDto } from './dto/response/channel-community-posts.dto';
import { ChannelCommunityPostsContinuationDto } from './dto/response/channel-community-posts-continuation.dto';
import { ChannelStatsDto } from './dto/response/channel-stats.dto';
import { ChannelInfoDto } from './dto/response/channel-info.dto';
import { SortType } from './types/sort';

@Injectable()
export class ChannelsService {
  constructor(
    @InjectModel(General.name)
    private readonly GeneralModel: Model<General>
  ) {}

  async getChannelHome(channelId: string): Promise<ChannelHomeDto> {
    if (!checkParams(channelId)) {
      throw new BadRequestException('Error fetching channel homepage', 'Invalid channelId');
    }
    try {
      return await ytch.getChannelHome({ channelId });
    } catch (error) {
      throwChannelError(error, 'Error fetching channel homepage');
    }
  }

  getChannelInfo(channelId: string): Promise<ChannelInfoDto> {
    if (!checkParams(channelId)) {
      throw new BadRequestException('Error fetching channel info', 'Invalid channelId');
    }
    try {
      return ytch.getChannelInfo({ channelId }) as unknown as Promise<ChannelInfoDto>;
    } catch (error) {
      throwChannelError(error, 'Error fetching channel info');
    }
  }

  async getChannelVideos(channelId: string, sort: SortType = 'newest'): Promise<ChannelVideosDto> {
    if (!checkParams(channelId, sort)) {
      throw new BadRequestException(
        'Error fetching channel videos',
        'Invalid channelId or sort parameter'
      );
    }
    try {
      return ytch.getChannelVideos({ channelId, sortBy: sort });
    } catch (error) {
      throwChannelError(error, 'Error fetching channel videos');
    }
  }

  getChannelVideosContinuation(continuation: string): Promise<ChannelVideosContinuationDto> {
    if (!checkParams(continuation)) {
      throw new BadRequestException('Error fetching channel videos', 'Invalid continuation string');
    }
    try {
      return ytch.getChannelVideosMore({ continuation });
    } catch (error) {
      throwChannelError(error, 'Error fetching channel videos');
    }
  }

  async getChannelShorts(channelId: string, sort = 'newest' as const): Promise<ChannelVideosDto> {
    if (!checkParams(channelId)) {
      throw new BadRequestException('Error fetching channel shorts', 'Invalid channelId');
    }
    try {
      return ytch.getChannelShorts({ channelId, sortBy: sort });
    } catch (error) {
      throwChannelError(error, 'Error fetching channel shorts');
    }
  }

  async getChannelLivestreams(
    channelId: string,
    sort = 'newest' as const
  ): Promise<ChannelVideosDto> {
    if (!checkParams(channelId)) {
      throw new BadRequestException('Error fetching channel livestreams', 'Invalid channelId');
    }
    try {
      return ytch.getChannelLivestreams({ channelId, sortBy: sort });
    } catch (error) {
      throwChannelError(error, 'Error fetching channel livestreams');
    }
  }

  getChannelPlaylists(channelId: string): Promise<ChannelPlaylistsDto> {
    if (!checkParams(channelId)) {
      throw new BadRequestException('Error fetching channel playlists', 'Invalid channelId');
    }
    try {
      return ytch.getChannelPlaylistInfo({ channelId });
    } catch (error) {
      throwChannelError(error, 'Error fetching channel playlists');
    }
  }

  getChannelPlaylistsContinuation(continuation: string): Promise<ChannelPlaylistsContinuationDto> {
    if (!checkParams(continuation)) {
      throw new BadRequestException(
        'Error fetching channel playlists',
        'Invalid continuation string'
      );
    }
    try {
      return ytch.getChannelPlaylistsMore({ continuation });
    } catch (error) {
      throwChannelError(error, 'Error fetching channel playlists');
    }
  }

  searchChannel(channelId: string, query: string): Promise<ChannelSearchDto> {
    if (!checkParams(channelId, query)) {
      throw new BadRequestException(
        'Error fetching channel channel search',
        'Invalid channelId or query'
      );
    }
    try {
      return ytch.searchChannel({ channelId, query });
    } catch (error) {
      throwChannelError(error, 'Error fetching channel channel search');
    }
  }

  searchChannelContinuation(continuation: string): Promise<ChannelSearchContinuationDto> {
    if (!checkParams(continuation)) {
      throw new BadRequestException('Error fetching channel search', 'Invalid continuation string');
    }
    try {
      return ytch.searchChannelMore({ continuation });
    } catch (error) {
      throwChannelError(error, 'Error fetching channel search');
    }
  }

  getRelatedChannelsContinuation(continuation: string): Promise<RelatedChannelsContinuationDto> {
    if (!checkParams(continuation)) {
      throw new BadRequestException(
        'Error fetching related channels',
        'Invalid continuation string'
      );
    }
    try {
      return ytch.getRelatedChannelsMore({ continuation });
    } catch (error) {
      throwChannelError(error, 'Error fetching related channels');
    }
  }

  getChannelCommunityPosts(channelId: string): Promise<ChannelCommunityPostsDto> {
    if (!checkParams(channelId)) {
      throw new BadRequestException('Error fetching channel community posts', 'Invalid channelId');
    }
    try {
      return ytch.getChannelCommunityPosts({ channelId });
    } catch (error) {
      throwChannelError(error, 'Error fetching channel community posts');
    }
  }

  getChannelCommunityPostsContinuation(
    continuation: string,
    innerTubeApi: string
  ): Promise<ChannelCommunityPostsContinuationDto> {
    if (!checkParams(continuation)) {
      throw new BadRequestException(
        'Error fetching channel community posts',
        'Invalid continuation string'
      );
    }
    try {
      return ytch.getChannelCommunityPostsMore({ continuation, innerTubeApi });
    } catch (error) {
      throwChannelError(error, 'Error fetching channel community posts');
    }
  }

  getChannelStats(channelId: string): Promise<ChannelStatsDto> {
    if (!checkParams(channelId)) {
      throw new BadRequestException('Error fetching channel stats', 'Invalid channelId');
    }
    try {
      return ytch.getChannelStats({ channelId });
    } catch (error) {
      throwChannelError(error, 'Error fetching channel stats');
    }
  }

  getTinyThumbnail(reply: FastifyReply, id: string): void {
    // eslint-disable-next-line dot-notation
    const imgPathWebp = path.join(global['__basedir'], `channels/${id}.webp`);
    // eslint-disable-next-line dot-notation
    const imgPathJpg = path.join(global['__basedir'], `channels/${id}.jpg`);

    const imageTransformer = sharp().resize(36, 36);

    try {
      const fileStream = fs.createReadStream(imgPathWebp);
      reply.type('image/webp').send(fileStream.pipe(imageTransformer));
      return;
    } catch {
      // Error is thrown later
    }

    try {
      const fileStream = fs.createReadStream(imgPathJpg);
      reply.type('image/jpeg').send(fileStream.pipe(imageTransformer));
      return;
    } catch {
      // Error is thrown later
    }

    throw new NotFoundException();
  }
}
