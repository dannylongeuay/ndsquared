FROM python:3.9-slim-buster as base

RUN apt-get update && apt-get install -y \
  tini \
  && rm -rf /var/lib/apt/lists/*

ENV HOME /usr/src
ENV PYTHONPATH $HOME
ENV VIRTUAL_ENV="$HOME/.venv"
ENV PATH="$VIRTUAL_ENV/bin:$PATH"

FROM base as builder

WORKDIR $HOME

COPY pyproject.toml poetry.lock ./

RUN pip install poetry && \
  poetry config virtualenvs.in-project true && \
  poetry install --no-dev

FROM base as app

RUN groupadd appuser && useradd -u 1000 -g appuser appuser

COPY --from=builder --chown=appuser:appuser $HOME/.venv/ $HOME/.venv
COPY --chown=appuser:appuser . $HOME

WORKDIR $HOME

ENTRYPOINT [ "/usr/bin/tini", "--" ]
