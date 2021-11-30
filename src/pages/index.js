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
        <Link to='https://github.com/mathbalduino/go-log/actions/workflows/ci.yml'>
          <img src="https://github.com/mathbalduino/go-log/actions/workflows/ci.yml/badge.svg?branch=main" />
        </Link>
        &nbsp; &nbsp;
        <Link to='https://codecov.io/gh/mathbalduino/go-log'>
          <img src="https://codecov.io/gh/mathbalduino/go-log/branch/main/graph/badge.svg?token=X6MM6EFXGT"/>
        </Link>
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
