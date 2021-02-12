import React, { useState } from 'react'
import Head from 'next/head'
import { useRouter } from 'next/router'

import styles from '../styles/Home.module.css'

export default function Home() {
  const [text, setText] = useState('')
  const [textError, setTextError] = useState(false)
  const [useTemplate, setUseTemplate] = useState(false)
  const router = useRouter()

  const submit = () => {
    if (text === '') {
      setTextError(true)
      return
    } else {
      router.push(`/image/${useTemplate ? 'time' : 'slide'}/${text.replace('\n', ',')}/share`)
    }
  }

  return (
    <div className={styles.container}>
      <Head>
        <title>換気タイム</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>
          さまざまな「タイム」を作成しよう！
        </h1>
        <form className={styles.forms}>
          <label>
            <span>{ textError ? 'お好みのタイムを入力してください' : '' }</span>
            <textarea className={styles.textArea} name="text" id="text" value={text} placeholder={'換気\nタイム'} onChange={(e) => setText(e.target.value)} />
          </label>
          <label>
            <input type="checkbox" name="useTemplate" value={useTemplate} onChange={() => setUseTemplate(!useTemplate)} />
            スライドのテンプレートを使う
          </label>
        </form>
        <a className={styles.submitBtn} onClick={submit}>
          確定
        </a>
      </main>
    </div>
  )
}
