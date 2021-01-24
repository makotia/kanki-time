import React, { useState } from 'react'
import Head from 'next/head'
import { useRouter } from 'next/router'

import styles from '../styles/Home.module.css'

export default function Home() {
  const [state, setState] = useState({
    useTemplate: false,
    line1: {
      text: '',
      isError: false,
    },
    line2: {
      text: '',
      isError: false,
    },
  })

  const router = useRouter()

  const setLine1 = (str) => {
    let s = state
    s.line1.text = str
    setState(s)
  }

  const setLine2 = (str) => {
    let s = state
    s.line2.text = str
    setState(s)
  }

  const submit = () => {
    if (state.line1.text === '') {
      let s = state
      s.line1.isError = true
      setState(!!s.line1.isError)
    }

    if (state.line2.text === '') {
      let s = state
      s.line2.isError = true
      setState(!state.line2.isError)
    }

    if (!(state.line1.isError || state.line2.isError)) {
      const body = {
        Text: [state.line1.text, state.line2.text].join('\n'),
        Type: state.useTemplate ? 'time' : 'slide',
      }
      fetch(`${process.env.apiUrl}/api`, {
        method: 'POST',
        body: JSON.stringify(body),
        headers: { 'Content-Type': 'application/json' },
      })
      .then(res => res.json())
      .then(json => router.push(`/${json.id}/share`))
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
            {state.line1.isError && <span>2行目が入力されていません</span>}
            <input type="text" name="name" value={state.textLine1} onChange={(e) => setLine1(e.target.value)} placeholder="換気" />
          </label>
          <label>
            {state.line2.isError && <span>2行目が入力されていません</span>}
            <input type="text" name="name" value={state.textLine2} onChange={(e) => setLine2(e.target.value)} placeholder="タイム" />
          </label>
          <label>
            <input type="checkbox" name="useTemplate" value={state.useTemplate} onChange={() => setUseTemplate(!state.useTemplate)} />
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
