import React from 'react'
import Head from 'next/head'

const Meta = ({ image, url }) => {
  return (
    <Head>
      <title>換気タイム</title>
      <link rel="icon" href="/favicon.ico" />
      <meta property="og:title" content="換気タイム" />
      <meta property="og:site_name" content="換気タイム" />
      <meta property="twitter:title" content="換気タイム" />
      <meta property="twitter:description" content="換気タイム" />
      <meta property="twitter:card" content="summary_large_image" />
      <meta property="twitter:site" content="@0x307E" />
      <meta property="twitter:url" content={url} />
      <meta property="og:url" content={url} />
      <meta property="og:description" content="換気タイム" />
      <meta property="og:image" content={image} />
      <meta property="twitter:image" content={image} />
    </Head>
  )
}

export default Meta
