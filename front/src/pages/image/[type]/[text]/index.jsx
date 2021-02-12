import React, { useEffect } from 'react'
import { useRouter } from 'next/router'
import Head from 'next/head'

export default function Home({ text, type }) {
  const url = process.env.baseUrl
  const router = useRouter()
  const imgURL = `${process.env.apiUrl}/api/?Text=${text}&Type=${type}`
  useEffect(() => router.push('/'))
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
        <meta property="og:image" content={imgURL} />
        <meta name="twitter:image" content={imgURL} />
      </Head>
    </div>
  )
}

export const getStaticProps = async (context) => {
  const { text, type } = context.params
  return {
    props: { text, type },
    revalidate: 100,
  }
}

export const getStaticPaths = async () => ({
  paths: [{params: {text: '', type: ''}}],
  fallback: 'blocking',
})
