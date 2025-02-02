import { useUserStore } from '@/store/user';

/**
 * This plugin runs user authentication server-side,
 * if there is an existing authentication cookie
 */
export default defineNuxtPlugin(async nuxtApp => {
  const userStore = useUserStore(nuxtApp.$pinia);

  const cookies = parseCookieString(nuxtApp.ssrContext.event.req.headers?.cookie);

  if (cookies?.Authentication) {
    await userStore.getUser(cookies.Authentication);
  }
});
