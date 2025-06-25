[homework questions link](https://github.com/DataTalksClub/llm-zoomcamp/blob/main/cohorts/2024/02-open-source/homework.md)

# Q1. Running Ollama with Docker
```bash
docker exec -it ollama bash
ollama -v
```
answer:
```
ollama version is 0.1.48
```

# Q2. Downloading an LLM
```bash
ollama pull gemma:2b
cd root/.ollama/models/manifests/registry.ollama.ai/library/
cat gemma/2b
```
answer:
```
{"schemaVersion":2,"mediaType":"application/vnd.docker.distribution.manifest.v2+json","config":{"mediaType":"application/vnd.docker.container.image.v1+json","digest":"sha256:887433b89a901c156f7e6944442f3c9e57f3c55d6ed52042cbb7303aea994290","size":483},"layers":[{"mediaType":"application/vnd.ollama.image.model","digest":"sha256:c1864a5eb19305c40519da12cc543519e48a0697ecd30e15d5ac228644957d12","size":1678447520},{"mediaType":"application/vnd.ollama.image.license","digest":"sha256:097a36493f718248845233af1d3fefe7a303f864fae13bc31a3a9704229378ca","size":8433},{"mediaType":"application/vnd.ollama.image.template","digest":"sha256:109037bec39c0becc8221222ae23557559bc594290945a2c4221ab4f303b8871","size":136},{"mediaType":"application/vnd.ollama.image.params","digest":"sha256:22a838ceb7fb22755a3b0ae9b4eadde629d19be1f651f73efb8c6b4e2cd0eea0","size":84}]}
```
# Q3. Running the LLM
```bash
ollama run gemma:2b
>>> 10*10
```
answer:
```
Sure, here's the answer to your question:

10 * 10 = 100.

I hope this is helpful!
```
# Q4. Donwloading the weights
```bash
du -h
```
answer:
```
1.6G    ./ollama_files/models/blobs
8.0K    ./ollama_files/models/manifests/registry.ollama.ai/library/gemma
12K     ./ollama_files/models/manifests/registry.ollama.ai/library
16K     ./ollama_files/models/manifests/registry.ollama.ai
20K     ./ollama_files/models/manifests
1.6G    ./ollama_files/models
1.6G    ./ollama_files
```
# Q5. Adding the weights
```Dockerfile
FROM ollama/ollama
COPY ./ollama_files /root/.ollama
```
# Q6. Serving it
```bash
docker build -t ollama-gemma2b -f week2/Dockerfile .
docker run -it --rm -p 11434:11434 ollama-gemma2b

pip install openai
python3.10 week2/chat-client.py
```

```python
from openai import OpenAI
import tiktoken

client = OpenAI(
    base_url='http://localhost:11434/v1/',
    api_key='ollama',
)

prompt = "What's the formula for energy?"

response = client.chat.completions.create(
    model='gemma:2b',
    messages=[{"role": "user", "content": prompt}],
    temperature=0.0
)

print(response.usage.completion_tokens)
```
link to the code: [python code](./chat-client.py)

answer:
```
304
```