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
  const { data } = useSWR(
    shouldFetch ? [`${process.env.apiUrl}/api/${id}`, ''] : null,
    fetcher
  )
  return (
    <div className={styles.container}>
      <Head>
        <title>Create Next App</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        {data &&
        <>
          <h1 className={styles.title}>
            { data }
          </h1>
        </>}
      </main>
    </div>
  )
}
