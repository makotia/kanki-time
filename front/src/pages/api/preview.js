export default async (req, res) => {
  if (!req.query.id) {
    return res.status(404).end();
  }

  res.setPreviewData({
    id: req.query.id,
  })
  res.writeHead(307, { Location: `/${req.query.id}` })
  res.end('Preview mode enabled')
};
