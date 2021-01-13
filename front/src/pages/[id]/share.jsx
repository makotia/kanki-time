import React, { useState, useEffect } from 'react'
import Head from 'next/head'

import useSWR from 'swr'
import { useRouter } from 'next/router'

import fetcher from '../../lib/fetcher'
import styles from '../../styles/Home.module.css'

export default function Home() {
  const router = useRouter()
  const [shouldFetch, setShouldFetch] = useState(false)
  const { id } = router.query
  useEffect(() => {
    setShouldFetch(true)
  }, [])
  const { data, error } = useSWR(
    shouldFetch ? [`${process.env.apiUrl}/api/${id}`, ''] : null,
    fetcher
  )
  if (data) this.props.router.push('/')
  return (
    <div className={styles.container}>
      <Head>
        <title>換気タイム</title>
        <meta property="og:title" content="換気タイム" />
        <link rel="icon" href="/favicon.ico" />
        <meta property="og:site_name" content="換気タイム" />
        <meta name="twitter:title" content="換気タイム" />
        <meta name="twitter:description" content="換気タイム" />
        <meta name="twitter:card" content="summary_large_image" />
        <meta name="twitter:site" content="@0x307E" />
        <meta name="twitter:url" content={process.env.baseUrl} />
        <meta property="og:url" content={process.env.baseUrl} />
        <meta property="og:description" content="換気タイム" />
        <meta property="og:image" content={(!error && id) ? `${process.env.apiUrl}/api/media/${id}.png` : ''} />
        <meta name="twitter:image" content={(!error && id) ? `${process.env.apiUrl}/api/media/${id}.png` : ''} />
      </Head>
    </div>
  )
}
