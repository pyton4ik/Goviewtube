export const parseCookieString = (str?: string | undefined | null): Record<string, string> =>
  str
    ?.split(';')
    ?.map(value => value.split('='))
    ?.reduce(
      (acc, value) => ({
        ...acc,
        [desanitize(value[0])]: desanitize(value[1])
      }),
      {}
    );
