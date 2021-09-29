import Layout from '@theme/Layout';
import clsx from 'clsx';
import React from 'react';
import styles from './index.module.css';

const Item = ({title, description, header}) => (
  <div className={clsx('col col--4')}>
    <div className="text--center" style={{display: 'flex', justifyContent: 'center', marginBottom: 20}}>
      {header}
    </div>
    <div className="text--center padding-horiz--md">
      <h3>{title}</h3>
      {description}
    </div>
  </div>
)

function Hello() {
  const BrasaoUfsc = require('../../static/img/brasao_ufsc.svg').default

  return (
    <Layout title="Hello">
      <div className={clsx('hero hero--primary', styles.heroBanner)}>
        <div className="container">
          <img style={{ marginBottom: '20px' }} src="https://dummyimage.com/200x200/8c8c8c/000000&text=Foto+perfil"/>
          <p className="hero__subtitle" style={{ marginBottom: 0 }}>
            Hey, be welcome! I'm <b>Matheus Leonel Balduino</b>, an independent brazilian
            developer that likes to create and publish open-source libraries, personal 
            projects, programming tips, tutorials, etc
          </p>
        </div>
      </div>
      <main>
        <section className={styles.features} style={{ margin: '30px 0' }}>
          <div className="container">
            <div className="row">
              <Item
                title='Where?' 
                description={
                  <p>
                    Currently, I'm working from <b>Florianópolis</b>, Santa Catarina (Brasil). What a beautiful city, don't you think?
                    <br />
                    Feel free to send me a message if I can help you with something
                  </p>
                } 
                header={
                  <iframe 
                    src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d452702.19589266065!2d-48.75047459416864!3d-27.57070558391982!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x9527394eb2c632d7%3A0x81bc550b6a04c746!2sFlorian%C3%B3polis%2C%20SC!5e0!3m2!1spt-BR!2sbr!4v1632502769176!5m2!1spt-BR!2sbr"
                    width="300"
                    height="300"
                    loading="lazy"
                  />
                } 
              />
              <Item
                title='Education' 
                description={
                  <p>
                    Computer Science at <a href="https://ufsc.br" style={{height: 300}} target="_blank">UFSC</a>,
                    that's what I'm doing. I'm almost finishing my graduation (if nothing goes wrong), but if 
                    you're having trouble finding me, the UFSC is a good starting point
                  </p>
                } 
                header={
                  <a href="https://ufsc.br" style={{height: 300}} target="_blank"><BrasaoUfsc width={300} height={300} /></a>
                } 
              />
              <Item
                title='Follow me via' 
                description={
                  <ul style={{listStyle: 'none', paddingLeft: 0}}>
                    <li><a href="https://github.com/mathbalduino" target="_blank">@mathbalduino</a> on GitHub</li>
                    <li><a href="https://instagram.com/mathbalduino" target="_blank">@mathbalduino</a> on Instagram</li>
                    <li><a href="https://twitter.com/mathbalduino" target="_blank">@mathbalduino</a> on Twitter</li>
                    <li><a href="https://youtube.com/c/mathbalduino" target="_blank">@mathbalduino</a> on YouTube</li>
                    <li><a href="https://www.twitch.tv/mathbalduino" target="_blank">@mathbalduino</a> on Twitch</li>
                    <li><a href="https://www.tiktok.com/@mathbalduino" target="_blank">@mathbalduino</a> on TikTok</li>
                    <li>...</li>
                  </ul>
                } 
                header={
                  <div style={{ width: 300, height: 300, display: 'flex', alignItems: 'center' }}>
                    <h1>That's all about @mathbalduino</h1>
                  </div>
                } 
              />
            </div>
          </div>
        </section>
      </main>
    </Layout>
  );
}

export default Hello;

/*
Hey! I'm Matheus Leonel Balduino, from Santa Catarina (Brasil), and I like to improve my developer skills building open-source libraries.

I'm  years old and currently studying at Universidade Federal de Santa Catarina (UFSC).

(Botar mapa, botar foto, botar algo da UFSC)

Feel free to ask me anything. You can find me everywhere using @mathbalduino. Try it out:
- [@mathbalduino](https://github.com/mathbalduino) on GitHub
- 
- [@mathbalduino]() on Twitter
- [@mathbalduino]() on YouTube

<div style={{ marginBottom: '30px', width: '100%', display: 'flex', justifyContent: 'space-around', alignItems: 'center', flexWrap: 'wrap' }}>
          <p style={{ maxWidth: '800px', textAlign: 'justify' }}>
          </p>
        </div>
        <div style={{ marginBottom: '30px', width: '100%', display: 'flex', justifyContent: 'space-around', alignItems: 'center', flexWrap: 'wrap' }}>
          <div style={{  width: '300px', height: '300px' }}>
            <iframe 
              src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d452702.19589266065!2d-48.75047459416864!3d-27.57070558391982!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x9527394eb2c632d7%3A0x81bc550b6a04c746!2sFlorian%C3%B3polis%2C%20SC!5e0!3m2!1spt-BR!2sbr!4v1632502769176!5m2!1spt-BR!2sbr"
              width="300"
              height="300"
              allowfullscreen=""
              loading="lazy"
            />
          </div>
          <p style={{ maxWidth: '800px', textAlign: 'justify' }}>
            Currently, I'm working from <b>Florianópolis</b>, Santa Catarina (Brasil). What a beautiful city, don't you think?
            <br /><br />
            Feel free to send me a message if I can help you with something
          </p>
        </div>
        <div style={{ marginBottom: '30px', width: '100%', display: 'flex', justifyContent: 'space-around', alignItems: 'center', flexWrap: 'wrap' }}>
          <p style={{ maxWidth: '800px', textAlign: 'justify' }}>
            
          </p>
          <div style={{  width: '300px', height: '300px' }}>
            <iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3535.7882560933776!2d-48.518724043417514!3d-27.600093605098312!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x952739005fe6a97f%3A0xbac58c3ef1d3885e!2sCentro%20Tecnol%C3%B3gico%20da%20UFSC!5e0!3m2!1spt-BR!2sbr!4v1632503010456!5m2!1spt-BR!2sbr" 
              width="300" 
              height="300" 
              allowfullscreen="" 
              loading="lazy"
            />
          </div>
        </div>
        <div style={{ width: '100%', display: 'flex', justifyContent: 'space-around', alignItems: 'center', flexWrap: 'wrap' }}>
          <p style={{ maxWidth: '800px', textAlign: 'justify' }}>
            
          </p>
          <div style={{ maxWidth: '350px', height: '300px',  display: 'flex', alignItems: 'center' }}>
            
          </div>
        </div>
 */
