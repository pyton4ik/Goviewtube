export const useProxyUrls = () => {
  const { apiUrl: clientApiUrl } = useApiUrl(true);
  return {
    textProxy: ``,
    imgProxy: ``,
    streamProxy: ``,
    videoPlaybackProxy: ``
  };
};
