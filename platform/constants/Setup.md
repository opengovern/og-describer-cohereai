# CohereAI Integration Guide

Learn how to integrate opencomply with your CohereAI project.

This guide shows you how to connect opencomply to a CohereAI project using a project-level API key with read-only access. Each integration corresponds to a single CohereAI project. Once integrated, opencomply will discover and assess key CohereAI resources—such as connectors, datasets, embedding jobs, fine-tuned models, and models—enabling compliance and visibility across your CohereAI environment.

## Prerequisites

- opencomply installed and running
- A CohereAI project with an appropriate read-only API key

## Create a Read-Only Project API Key

1. Log in to the CohereAI Dashboard and select the relevant project.
2. Go to **Settings** or **Manage Keys** (the exact location may vary).
3. Click **Create New Key**, ensure it is **Read-Only**, and copy the generated key.

## Configure Integration in opencomply

1. In the opencomply dashboard, go to **Integrations > CohereAI**.
2. Select **API Key integration**.
3. Paste in your CohereAI project-level read-only API key.
4. Specify the CohereAI project name.
5. Click **Complete** to govern.
6. Click **Save**.