export const fetcher = (url) =>
  fetch(url, {
    method: 'GET',
  }).then((r) => {
    if (!r.ok) {
      throw new Error(String(r.status))
    } else if (r.body != '') {
      return r
    }
    return r.json()
  })

export default fetcher
