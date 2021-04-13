import React from 'react'
import Head from 'next/head'

type Props = {
  image: string;
  url: string;
}

const Meta: React.FC<Props> = ({ image, url }) => {
  return (
    <Head>
      <title>換気タイム</title>
      <link rel="icon" href="/favicon.ico" />
      <meta property="og:title" content="換気タイム" />
      <meta property="og:site_name" content="換気タイム" />
      <meta name="twitter:title" content="換気タイム" />
      <meta name="twitter:description" content="換気タイム" />
      <meta name="twitter:card" content="summary_large_image" />
      <meta name="twitter:site" content="@0x307E" />
      <meta name="twitter:url" content={url} />
      <meta property="og:url" content={url} />
      <meta property="og:description" content="換気タイム" />
      <meta property="og:image" content={image} />
      <meta name="twitter:image" content={image} />
    </Head>
  )
}

export default Meta
