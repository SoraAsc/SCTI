import os
import requests
import json

github_token = os.getenv('GITHUB_TOKEN')
trello_key = os.getenv('TRELLO_KEY')
trello_token = os.getenv('TRELLO_TOKEN')
trello_list_id = os.getenv('TRELLO_LIST_ID')
github_repo = 'MintzyG/SCTI'

def get_github_prs(repo):
    url = f'https://api.github.com/repos/{repo}/pulls'
    headers = {
        'Authorization': f'token {github_token}',
        'Accept': 'application/vnd.github.v3+json'
    }
    response = requests.get(url, headers=headers)
    response.raise_for_status()
    return response.json()

def get_trello_cards():
    url = f'https://api.trello.com/1/lists/{trello_list_id}/cards'
    headers = {
        'Accept': 'application/json'
    }
    query = {
        'key': trello_key,
        'token': trello_token
    }
    print(f'Trello API Request URL: {url}')
    print(f'Trello API Query Parameters: {query}')
    response = requests.get(url, headers=headers, params=query)
    response.raise_for_status()
    
    try:
        return response.json()
    except json.JSONDecodeError:
        print("Failed to decode JSON response from Trello.")
        return []

def create_trello_card(name, desc):
    url = 'https://api.trello.com/1/cards'
    query = {
        'idList': trello_list_id,
        'name': name,
        'desc': desc,
        'key': trello_key,
        'token': trello_token
    }
    response = requests.post(url, params=query)
    response.raise_for_status()
    return response.json()

def main():
    prs = get_github_prs(github_repo)
    
    existing_cards = get_trello_cards()
    existing_card_names = {card['name'] for card in existing_cards}
    
    for pr in prs:
        pr_name = pr['title']
        pr_url = pr['html_url']
        
        if pr_name not in existing_card_names:
            created_card = create_trello_card(pr_name, pr_url)
            print(f'Created Trello card for PR: {pr_name}')
            print(json.dumps(created_card, sort_keys=True, indent=4, separators=(",", ": ")))
        else:
            print(f'Trello card for PR already exists: {pr_name}')

if __name__ == "__main__":
    main()
