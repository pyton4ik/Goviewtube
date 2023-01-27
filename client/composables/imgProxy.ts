export const useImgProxy = () => {
  const { imgProxy } = useProxyUrls();

  const proxyUrl = (url: string) => {
    let correctedUrl = url;
    if (!url) return '';
    if (url.startsWith('//')) {
      correctedUrl = `https:${url}`;
    }
    return `${correctedUrl}`;
  };

  return { url: imgProxy, proxyUrl };
};
