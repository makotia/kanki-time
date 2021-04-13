export const fetcher = (url: string) =>
  fetch(url, {
    method: 'GET',
  }).then((r) => {
    if (!r.ok) {
      throw new Error(String(r.status))
    } else if (r.body != null) {
      return r
    }
    return r.json()
  })

export default fetcher
