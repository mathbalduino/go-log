import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Link from '@docusaurus/Link';
import Layout from '@theme/Layout';
import clsx from 'clsx';
import React from 'react';
import HomepageFeatures from '../components/HomepageFeatures';
import styles from './index.module.css';

function HomepageHeader() {
  const {siteConfig} = useDocusaurusContext();
  const mathbalduino_logoM = require('../../static/img/mathbalduino_logoM.png').default

  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className="container">
        <img src={mathbalduino_logoM} />
        <h1 className="hero__title">{siteConfig.title}</h1>
        <p className="hero__subtitle">{siteConfig.tagline}</p>
        <div className={styles.buttons} style={{ marginBottom: 'var(--ifm-paragraph-margin-bottom)' }}>
          <Link
            className="button button--secondary button--lg"
            to="/docs/intro">
            Read the Introduction 📜
          </Link>
        </div>
        <img src="https://github.com/mathbalduino/go-log/actions/workflows/go.yml/badge.svg?branch=main" />
      </div>
    </header>
  );
}

export default function Home() {
  return (
    <Layout
      description="go-log library documentation by @mathbalduino">
      <HomepageHeader />
      <main>
        <HomepageFeatures />
      </main>
    </Layout>
  );
}
