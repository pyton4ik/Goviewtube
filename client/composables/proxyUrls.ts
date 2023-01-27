export const useProxyUrls = () => {
  const { apiUrl: clientApiUrl } = useApiUrl(false);
  return {
    textProxy: ``,
    imgProxy: ``,
    streamProxy: ``,
    videoPlaybackProxy: ``
  };
};
