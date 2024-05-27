/** @type {import('next').NextConfig} */
module.exports = {
  reactStrictMode: true,
  transpilePackages: ["@repo/ui"],
  async redirects() {
    return [
      {
        source: '/docs',
        destination: 'http://docs.dotem.website',
        permanent: true,
      },
    ]
  },
};
