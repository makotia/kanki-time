import React, { useEffect } from 'react'
import { useRouter } from 'next/router'
import Head from 'next/head'

export default function Home({ id }) {
  const url = process.env.baseUrl
  console.log(id)
  return (
    <div>
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
        <meta property="og:image" content={`${process.env.apiUrl}/api/media/${id}.png`} />
        <meta name="twitter:image" content={`${process.env.apiUrl}/api/media/${id}.png`} />
      </Head>
    </div>
  )
}

export const getStaticProps = async (context) => {
  const id = context.params.id
  return {
    props: { id },
    revalidate: 1,
  }
}

export const getStaticPaths = async () => ({
  paths: [{params: {id: '1'}}],
  fallback: true,
})
