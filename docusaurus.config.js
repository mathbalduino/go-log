const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

// With JSDoc @type annotations, IDEs can provide config autocompletion
/** @type {import('@docusaurus/types').DocusaurusConfig} */
(module.exports = {
  title: 'go-log',
  tagline: 'everywhere, under @mathbalduino',
  url: 'https://mathbalduino.github.io',
  baseUrl: '/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',
  organizationName: 'mathbalduino', // Usually your GitHub org/user name.
  projectName: 'go-log', // Usually your repo name.
  trailingSlash: false,

  presets: [
    [
      '@docusaurus/preset-classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          editUrl: 'https://github.com/mathbalduino/go-log/edit/main/docs/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      navbar: {
        title: 'go-log',
        logo: {
          alt: '@mathbalduino logo',
          src: 'img/mathbalduino_logoM.png',
        },
        items: [
          {
            type: 'doc',
            docId: 'intro',
            position: 'left',
            label: 'Documentation',
          },
          {
            to: 'about',
            position: 'left',
            label: 'About',
          },
          {
            href: 'https://github.com/mathbalduino/go-log',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          {
            title: 'Documentation',
            items: [
              {
                label: 'Introduction',
                to: '/docs/intro',
              },
              {
                label: 'Basics',
                to: '/docs/basic-concepts/log_fields',
              },
              {
                label: 'Advanced',
                to: '/docs/advanced/loggers_clonage',
              },
            ],
          },
          {
            title: 'Author',
            items: [
              {
                label: 'mathbalduino.com.br',
                href: 'https://mathbalduino.com.br',
              },
            ],
          },
        ],
        logo: {
          alt: '@mathbalduino logo',
          href: 'http://mathbalduino.com.br',
          src: 'img/mathbalduino_logoS.png'
        },
        copyright: '@mathbalduino (Built with Docusaurus)',
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
    }),
});
